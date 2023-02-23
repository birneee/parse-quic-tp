# QUIC Transport Parameter Parser CLI

## Example 

```bash
$ echo "BQSACAAABgSACAAABwSACAAABASADAAACAJAZAkCQGQBAlOIAwJFrAsBGgIQAY5EgccmfdVjFdhWb5oC0gAKsBDB/4fAmc2Meg4BBA8ELt4auA==" | base64 -d | go run .
initial_max_stream_data_bidi_local, 524288
initial_max_stream_data_bidi_remote, 524288
initial_max_stream_data_uni, 524288
initial_max_data, 786432
initial_max_streams_bidi, 100
initial_max_streams_uni, 100
max_idle_timeout, 5000
max_udp_payload_size, 1452
max_ack_delay, 26
stateless_reset_token, [1 142 68 129 199 38 125 213 99 21 216 86 111 154 2 210]
original_destination_connection_id, [176 16 193 255 135 192 153 205 140 122]
active_connection_id_limit, 4
initial_source_connection_id, [46 222 26 184]
```
