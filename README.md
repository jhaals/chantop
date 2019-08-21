# chantop

chantop queries the gRPC [channelz service](https://github.com/grpc/proposal/blob/master/A14-channelz.md) for statistics about open channels.

```console
$ chantop -address localhost:8080
+---------------------------------------------------------------------------+-------+---------+-----------+--------+--------------------------+
|                                  TARGET                                   | STATE | STARTED | SUCCEEDED | FAILED |        LAST CALL         |
+---------------------------------------------------------------------------+-------+---------+-----------+--------+--------------------------+
| dns://service-name                                                        | READY |     247 |       247 |      0 | 2019-08-21T13:39:35.696Z |
| [[[some-service131.service.exampledomain.com./10.182.40.35:8080]/{}]]     | READY |      16 |        16 |      0 | 2019-08-21T13:39:34.655Z |
| [[[some-service13200.service.exampledomain.com./10.182.152.123:8080]/{}]] | READY |     106 |       106 |      0 | 2019-08-21T13:39:35.696Z |
| [[[some-service113.service.exampledomain.com./10.182.37.22:8080]/{}]]     | READY |      15 |        15 |      0 | 2019-08-21T13:39:33.613Z |
+---------------------------------------------------------------------------+-------+---------+-----------+--------+--------------------------+
```
