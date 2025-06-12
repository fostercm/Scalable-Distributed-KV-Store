# Scalable-Distributed-KV-Store


## Goals
- [Ease of use](#ease-of-use)
- [Scalability](#scalability)
- [Fault tolerance](#fault-tolerance)
- [Observability](#observability)
- [Efficiency](#efficiency)
- [Security](#security)

### Ease of use
- Simple user interface (reminiscent of a standard Go map)
- Straightforward background setup
- Abstracted routing, retry logic, and error handling

### Scalability
- Unlimited horizontal scaling
- Dynamic upscaling/downscaling
- Minimal performance decrease at immense scale

### Fault tolerance
- High availability
- Server failures do not result in unreachable/deleted keys
- No single points of failure

### Observability
- Simple monitoring
- Integration with Prometheus/Grafana
- Useful metrics and logging to understand system internals

### Efficiency
- No wasted compute
- Memory-safety
- Early failures

### Security
- Storage access secured
- Layers of protection against data leakage
- Authentication for network components

## Phases
This project was implemented in multiple stages, each with a specific goal in order to show progression in knowledge and ability:
- Phase 0: [Problem definition and planning](https://github.com/fostercm/Scalable-Distributed-KV-Store)
- Phase 1: [Basic functionality](https://github.com/fostercm/Scalable-Distributed-KV-Store/tree/Phase-1-Basic-Functionality)
- Phase 2: [Sharding and scaling](https://github.com/fostercm/Scalable-Distributed-KV-Store/tree/Phase-2-Sharding/Scaling)
- Phase 3: Dynamic scaling
- Phase 4: Fault tolerance
- Phase 5: Observability
- Phase 6: Security
- Phase 7: Testing and polishing
