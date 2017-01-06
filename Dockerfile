FROM northtyphoon/beego:latest

MAINTAINER Bin Du <northtyphoon@gmail.com>

# Install app dependencies
RUN mkdir -p "$GOPATH/src/github.com/northtyphoon/ballpen"
WORKDIR "$GOPATH/src/github.com/northtyphoon/ballpen"
COPY package.json "$GOPATH/src/github.com/northtyphoon/ballpen"
RUN npm install

# Deploy app source
COPY . "$GOPATH/src/github.com/northtyphoon/ballpen"
RUN npm run build-prod

# App port
EXPOSE 8080

# Admin port
EXPOSE 8088

CMD [ "bee", "run" ]