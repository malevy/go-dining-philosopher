# Dining Philosophers Problem

A classic problem from computer science for demonstrating concurrency
and synchronization issues.

There are K philosophers living together. When it's time for dinner,
they all sit at the same table. Between each philosopher is a fork.
In order to eat, a philosopher must pick up the fork on either side. Both 
forks must be acquired in order to eat. When a philosopher done eating, 
the forks are placed back down on the table.

Now, Philosophers being philosophers, they never stop thinking, 
even at dinner. So the typical pattern is that a Philosopher will think
for some period of time, become hungry, pick up the fork on both side and 
begin eating. When the hunger subsides, the forks are put down and the 
philosopher begins thinking again.

### Goal
Design a protocol that allows the philosophers to enjoy their meal without
causing a deadlock. If a deadlock occurs, all the philosophers will starve
to death.

