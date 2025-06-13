# Scalable-Distributed-KV-Store
I am in the process of designing a distributed key-value store in Go inspired by Apache Cassandra and coursework from the MIT graduate course **6.5840: Distributed Systems**.
The overall goal of this project is to get an intimate understanding of distributed systems fundamentals, challenges, and techniques at scale.

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
- Phase 0: [Problem definition and planning](#phase-0-problem-definition-and-planning)
- Phase 1: [Basic functionality](https://github.com/fostercm/Scalable-Distributed-KV-Store/tree/Phase-1-Basic-Functionality)
- Phase 2: [Sharding and scaling](https://github.com/fostercm/Scalable-Distributed-KV-Store/tree/Phase-2-Sharding/Scaling)
- Phase 3: Dynamic scaling
- Phase 4: Fault tolerance
- Phase 5: Observability
- Phase 6: Security
- Phase 7: Testing and polishing

## Collaboration
This project is a solo endeavor I am taking on to master distributed systems concepts so pull requests will not be reviewed.
However, if you stumble across this work and would like to give suggestions or pointers I have created a **suggestion** issue label.

## Acknowledgements
Thank you to the CS department at MIT for making your course materials public and thus inspiring this project.
That being said, all code in this project is written by myself and not taken from any course materials.

## License
This project is licensed under the MIT license, check the LICENSE file to learn more.

# Phase 0: Problem definition and planning
