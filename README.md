[![Maintainability](https://api.codeclimate.com/v1/badges/3f4a76f5228da102893f/maintainability)](https://codeclimate.com/github/neomede/chattp2/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/3f4a76f5228da102893f/test_coverage)](https://codeclimate.com/github/neomede/chattp2/test_coverage)

# chattp2

chattp2 is just a small chat application to play with golang and HTTP2.

### Usage

- Build the application

```sh
make build
```

- Launch chattp2 server:

```sh
./chattp2
```

- Launch clients:

Need to pass the sender name and the receiver name as a parameters

```sh
./chattp2 --client --sender=sender_name --receiver=receiver_name
```

### TODO

- Allow send messages to different users.
- Improve CLI. Printing messages and reading input.
- Add chat rooms.

### Resources

- HTTP2 [demo](https://github.com/golang/net/blob/master/http2/h2demo/h2demo.go) from http/net
- [High Performance Browser Networking](https://hpbn.co/http2/)
- [Example](https://github.com/golang/go/issues/13444) for bidirectional HTTP2 Streams

### Developed By

Rubén Simón Andreo - <rsimonandreo@gmail.com>

<a href="https://twitter.com/neomede">
  <img alt="Twitter" src="https://g.twimg.com/Twitter_logo_blue.png" width="100px"/>
</a>

<a href="https://es.linkedin.com/in/rubensimon">
  <img alt="Linkedin" src="https://media.licdn.com/mpr/mpr/shrink_200_200/AAEAAQAAAAAAAANyAAAAJGRlZTNlZDQwLTk4YTItNDA1MS04MzBjLWJmNGQ5M2RmZGUxYw.png" width="100px"/>
</a>

License
-------

    Copyright 2016 Rubén Simón Andreo

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
