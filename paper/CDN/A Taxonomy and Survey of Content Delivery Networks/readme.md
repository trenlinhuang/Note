# A Taxonomy and Survey of Content Delivery Networks
source: http://www.hit.bme.hu/~jakab/edu/litr/CDN/CDN-Taxonomy.pdf
## Abstract
**purpose**: Our aim is to categorize and analyze the existing CDNs, and to explore the uniqueness, weaknesses, opportunities, and future directions in this field.

**In this paper**
* provide a comprehensive taxonomy with a broad coverage of CDNs in terms of
  * organizational structure
  * content distribution mechanisms
  * request redirection techniques
  * performance measurement methodologies
* study the existing CDNs in terms of
  * their infrastructure
  * request-routing mechanisms
  * content replication techniques
  * load balancing
  * and cache management
* provide an indepth analysis and **state-of-the-art** survey of CDNs

  > state-of-the-art: 最先进的
*  apply the taxonomy to map various CDNs<br> *The mapping of the taxonomy to the CDNs helps in "gap" analysis in the content networking domain. It also provides a means to identify the present and future development in this field and validates the applicability and accuracy of the taxonomy.*

## 1. Introduction
### problem
popular Web services often suffer congestion and bottlenecks due to large demands made on their services.

**consequence**
*  unmanageable levels of traffic flow
*  many requests being lost

### How CDN deal with it
Content Delivery Networks (CDNs) [8][16][19] provide services that improve network performance by
* maximizing bandwidth
* improving accessibility
* maintaining correctness

**through content replication**.

They offer fast and reliable applications and services by distributing content to cache or edge servers located close to users [8].

### A CDN has some combination of
* `content-delivery infrastructure`
  * consists of a set of edge servers (also called surrogates) that deliver copies of content to end-users
  *  moves content from the origin server to the CDN edge servers and ensures consistency of content in the caches
* `request-routing distribution`  is responsible to directing client request to appropriate edge servers.
* `accounting infrastructure` maintains logs of client accesses and records the usage of the CDN servers. This info rmation is used for traffic reporting and usage-based billing.

### What CDNs need to do
A CDN focuses on building its network infrastructure to provide the following services and functionalities:
* storage and management of content
* distribution of content among surrogates
* cache management
* delivery of static
* dynamic and streaming content
* backup and disaster recovery solutions
* and monitoring
* performance measurement and reporting.

### Our contributions
A few studies have investigated CDNs in the recent past.

...

**none of these works has categorized CDNs**, in this work we focus on developing a taxonomy and presenting a detailed survey of CDNs.

### Key contributions
1. Develop a comprehensive taxonomy of CDNs
2. 