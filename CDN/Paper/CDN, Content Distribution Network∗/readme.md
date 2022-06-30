# CDN: Content Distribution Network∗
source: https://arxiv.org/pdf/cs/0411069.pdf
## Abstract
**Problem:**
the available network bandwidth and server capacity continue to be overwhelmed by the skyrocketing Internet utilization and the accelerating growth of bandwidth intensive content.

**How CDN Works:**
CDN replicates the content from the place of origin to the replica servers scattered over the Internet and serves a request from a replica server close to where the request originates.

**In This Paper:**
* give an overview about CDN.
* present the critical issues involved in designing and implementing an effective CDN
* survey the approaches proposed in literature to address these problems
* present a scheme that provides fast service location for peer-to-peer systems, a special type of CDN **with no infrastructure support.❓**
>**with no infrastructure support** 根据后文看意思是指不用通过增加硬件设备、带宽等来达到目的。<br>**"with infrastructure support":**
One might think that the constant improvement in the bandwidth of Internet infrastructure, for example, the availability of high-speed "last mile" connection of the subscribers to the Internet and the backbone fibers, and the increasing capacity of the various servers would reduce or eliminate the access delay problem eventually.<br>
from: Introduction

## Introduction
### Problem
Although the time to load the content of key Internet sites has improved constantly over the last several years, overall Web content access latency is still in the range of a few seconds, which is several times the threshold believed to represent natural human reading/scanning speeds.（资源下载速度还是太慢了）

### Why The Problem come into existence
* lack of overall management for Internet
  *  the absence of overall administration in turn makes it very difficult to guarantee proper performance and deal systematically with performance problems.
* as the load on Internet and the richness of its content continue to soar, any increase in available network bandwidth and server capacity will continue to be overwhelmed.

These two facts not only elevate the delay in accessing content on Internet, but also make the access latency unpredictable.

### How to deal with it, The CDN Way
The basic approach to address the performance problem is to move the content from the places of origin servers to the places at the edge of the Internet.

**Benefits**
* has better performance (lower access latency, higher transfer rate) than from the origin server
* using multiple replica servers for servicing requests costs less than using only the **data communications network❓** does

  > Data Communications: Data transmission and data reception or, more broadly, **data communication** or digital communications is the transfer and reception of data in the form of a digital bitstream or a digitized analog signal[1] over a point-to-point or point-to-multipoint communication channel.<br>from: https://en.wikipedia.org/wiki/Data_communication

**How Dose The CDN Do:**
CDN replicates a very selective set of content to the replica servers and only sends those requests for the replicated content to a replica server.

**Key Challenges:**
* How to place replica servers
* distribute content copies to replica servers
* how to route the requests to the proper replica server having the desired content 

## Overview
...
the fact that many objects are **not cacheable but replicable❓**, which include dynamic objects with read-only access and personalized objects (e.g., "cookied" requests), makes CDN indispensable.

![A General Architecture of CDN](./System%20Architecture%20Components%20of%20a%20CDN.png)<br>
Figure 1: System Architecture Components of a CDN
### 2.1 A General Architecture of CDN
