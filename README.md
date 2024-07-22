# Server-Sent Events (SSE) proxy

Proof of concept for a Server-Sent Events (SSE) proxy server.

The target server codes were taken from the [Go Echo example repo](https://github.com/labstack/echox).

## About SSE protocol

The proper SSE (Server-Sent Events) formatting is essential for the correct functioning of the protocol.
The following is a detailed explanation of why this formatting is necessary.

### Protocol Specifications

SSE is a standardized protocol defined by the HTML5 specification,
which outlines how messages should be formatted for compatibility across different clients (typically web browsers).

### Message Delimitation

`data:` **Prefix**: Each line of data in an SSE message must be prefixed with `data:`.
This tells the client that the line contains data to be processed as part of an event.
Without this prefix, the client wouldn't recognize the data correctly.

`double newline (\\n\\n)`: A double newline sequence (`\n\n`) indicates the end of an event.
This ensures that the client knows when one event ends and another begins.
It allows the client to process the entire event before moving on to the next one.

### Event Fields

The SSE protocol allows for several fields in a message:

* `data:`: The main data field, which can appear multiple times.
* `event:`: Specifies the event type.
* `id:`: Sets the event ID, which the client can use to resume from the last event.
* `retry:`: Suggests the reconnection time in milliseconds.

Proper formatting ensures that the client will parse these fields correctly.

### Client Parsing Logic

Web browsers and other SSE clients have built-in parsers that expect this specific format.
If the data is not formatted correctly, the client will not be able to parse the events correctly,
leading to issues like missing events, incorrect data interpretation, or connection errors.

### Example of SSE Message

```html
data: message part 1
data: message part 2
event: customEvent
id: 12345

data: another message
```

The above example shows a multi-line data event, a custom event type and an event ID.
The double newline separates this event from the next one.

By adhering to these formatting rules, you ensure that the SSE connection functions correctly 
and the client can reliably parse and process the incoming events.