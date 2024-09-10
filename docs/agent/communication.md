# Agent communication

## MQTT

```mermaid
sequenceDiagram
    Agent->>Server: REST: Request auth token
    Server->>Agent: REST: Response
    Agent->>Server: MQTT: Auth with token
    Agent-->>Server: MQTT: Subscribe to topics
```

### Topics

* `pixconf/agent/<agent name>/commands` - Read: Agent listens for commands from the server
* `pixconf/agent/<agent name>/response/<uuid>` - Write: Agent sends response of a command to the server
* `pixconf/agent/<agent name>/health` - Write: Agent sends health status to the server
