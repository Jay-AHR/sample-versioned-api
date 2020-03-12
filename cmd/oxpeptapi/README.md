To build `registry.gitlab.com/omt-mdl/peptide/oxpeptapi:latest`, issue the following command sequence in this directory:

1. `docker login registry.gitlab.com` - uses your Gitlab creds, same as for git operations
2. `docker build -t registry.gitlab.com/omt-mdl/peptide/oxpeptapi:latest .`
3. `docker push registry.gitlab.com/omt-mdl/peptide/oxpeptapi:latest`

To run `registry.gitlab.com/omt-mdl/peptide/oxpeptapi:latest`, issue the following sequence of commands:

1. `docker pull registry.gitlab.com/omt-mdl/peptide/oxpeptapi:latest`
2. `docker run -p 8080:8080 registry.gitlab.com/omt-mdl/peptide/oxpeptapi:latest` (please note that port 8080 is exposed in the container and this command binds to host port 8080)

To test, issue the following on the same machine on which the container is running (all possible options are shown following both `lifecycle` and `chaincode` in the URLs below):

1. curl http://localhost:8080/v2/peer/lifecycle/package|install|queryinstalled|getinstalledpackage|approveformyorg|checkcommitreadiness|commit|querycommitted or curl http://localhost:8080/v2/peer/chaincode/install|instantiate|invoke|package|query|signpackage|upgrade|list