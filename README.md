# Readers-Writers problem
Readers-writers problem portrays a common situation where one piece of data is both read and modified. There are two types of processes: readers and writers. Multiple readers can read data at the same time. Writer needs exclusive access to data, to modify it. The challenge in this problem is establishing a priority for writers in such a way that readers are not starved.

## Solution
Problem that is portrayed by readers-writers problem is so common that modern languages provide a ready to use abstraction. In the case of Go it is a RWMutex that is a recommended mechanism for solving that problem.

Four methods of RWMutex will be used in the solution:
1. Lock - locks RWMutex for writing.
2. Unlock - unlocks RWMutex for writing. 
3. RLock - locks RWMutex for reading.
4. RUnlock - unlocks RWMutex for reading.

Method Lock and Unlock should be used for critical section of the writers function. RLock and RUnlock should be used for readers critical section. Priority of writers is achieved by blocking new readers if a writer process is waiting. Writer waits for all reader processes to finish reading. After writer modified the data it first unblocks reader processes. Readers have a chance to access data before writers, and therefore will not be starved.

## Wikipedia
https://en.wikipedia.org/wiki/Readers%E2%80%93writers_problem