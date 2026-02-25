# Ch-1 Reliable, Scalable, and Maintainable Applications

## Building blocks of a data-intensive application
- Database: stores data so it can be found later
- Cache: remembers the result of an expensive operation, to speed up reads
- Search index: allows users to search data by keyword or filter it in various ways
- Stream processing: used to send a message to another process, to be handled asynchronously
- Batch processing: used to periodically crunch a large amount of accumulated data

## Data system
Combines several smaller general-purpose components like cache, database, message queue, search engines etc to create a new special-purpose data system

## Concerns of software systems
- Reliability: The system should continue to work correctly even in the face of hardware or software faults, and human errors
- Scalability: As system grows in data volume, traffic volume or complexity, there should be reasonable ways of dealing with that growth
- Maintainability: System should be maintainable for people working on it in the future

Operation- O(n^2)

dataset = 100,000

Single store with 100k entries
- Cost = 100,000,000,000

2 stores with 50k entries
- Cost = 50k^2 * 2 = 5,000,000,000

10 stores with 10k entries
- Cost = 10k^2 * 10 = 1,000,000,000
