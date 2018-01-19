package server

import (
	"context"
	"math/big"
	"net/http"

	"time"

	"github.com/alpe/community-bifrost/pkg/bitcoin"
	"github.com/alpe/community-bifrost/pkg/config"
	"github.com/alpe/community-bifrost/pkg/database"
	"github.com/alpe/community-bifrost/pkg/ethereum"
	"github.com/alpe/community-bifrost/pkg/queue"
	"github.com/alpe/community-bifrost/pkg/sse"
	"github.com/alpe/community-bifrost/pkg/stellar"
	"github.com/stellar/go/support/log"
)

// ProtocolVersion is the version of the protocol that Bifrost server and
// JS SDK use to communicate.
const ProtocolVersion int = 1

type Server struct {
	BitcoinListener            *bitcoin.Listener            `inject:""`
	BitcoinAddressGenerator    *bitcoin.AddressGenerator    `inject:""`
	Config                     *config.Config               `inject:""`
	Database                   database.Database            `inject:""`
	EthereumListener           *ethereum.Listener           `inject:""`
	EthereumAddressGenerator   *ethereum.AddressGenerator   `inject:""`
	StellarAccountConfigurator *stellar.AccountConfigurator `inject:""`
	TransactionsQueue          queue.Queue                  `inject:""`
	SSEServer                  sse.ServerInterface          `inject:""`

	MinimumValueBtc string
	MinimumValueEth string

	minimumValueSat            int64
	minimumValueWei            *big.Int
	httpServer                 *http.Server
	log                        *log.Entry
	stopTransactionQueueWorker context.CancelFunc
	queueRetryDelay            time.Duration
}

type GenerateAddressResponse struct {
	ProtocolVersion int    `json:"protocol_version"`
	Chain           string `json:"chain"`
	Address         string `json:"address"`
}
