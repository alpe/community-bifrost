package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/asaskevich/govalidator"
	configsupport "github.com/stellar/go/support/config"
	"github.com/stellar/go/support/errors"
)

func Decode(content string, dest interface{}) error {
	metadata, err := toml.Decode(content, dest)
	if err != nil {
		return errors.Wrap(err, "decode-file failed")
	}

	// Undecoded keys correspond to keys in the TOML document
	// that do not have a concrete type in config struct.
	undecoded := metadata.Undecoded()
	if len(undecoded) > 0 {
		return errors.New("Unknown fields: " + fmt.Sprintf("%+v", undecoded))
	}

	valid, err := govalidator.ValidateStruct(dest)

	if valid {
		return nil
	}

	fields := govalidator.ErrorsByField(err)

	return &configsupport.InvalidConfigError{
		InvalidFields: fields,
	}
}
