## ci folder 
contains build related artifacts for docker Content Trust signed images.

## Resources
* Docker content trust: https://docs.docker.com/engine/security/trust/content_trust
* Notary: see https://github.com/theupdateframework/notary
### Add
```bash

alias dockernotary="$GOPATH/bin/notary -s https://notary.docker.io -d ~/.docker/trust"

# Setup notary project
dockernotary init -p alpetest/mybifrost

# Rotate any existing key
dockernotary key rotate docker.io/alpetest/mybifrost snapshot -r

# add private key to docker trust folder `~/.docker/trust/private`
dockernotary key import delegation.key --role user

# Add new delegation cert to project
dockernotary delegation add docker.io/alpetest/mybifrost targets/releases delegation.crt --all-paths
# publish changes
dockernotary publish docker.io/alpetest/mybifrost

# check new key is in delegation key list
dockernotary delegation list docker.io/alpetest/mybifrost
# check targets in notary
dockernotary list docker.io/alpetest/mybifrost
```

### Remove target
```bash
dockernotary remove docker.io/alpetest/mybifrost manual --publish
```


