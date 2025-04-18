# go-data-structures
Data structures in Golang.

With this module, I am to (eventually) create common data structures, such as the linked list, doubly linked list, circular list, stack, queue, and so on, in Go.

# LinkedList data structure
The LinkedList data structure is used to create a list of integer objects (as of now, eventually hope to make it more type-independent).

# LinkedList Current Methods
Push() - pushes an integer, passed as an argument, onto the list at the tail.

Pop() - removes and returns the tail object on the linked list, updates tail to point to the previous object in the list

Traverse() - prints the objects in the list, from head to tail

# LinkedList Potential Methods
View() - returns the tail object on the linked list without removing it

Poke() - removes and returns the head object on the linked list, updates head to point to the next object in the list

Peep() - returns the head object on the linked list without removing it

Insert() - inserts an object, passed as an argument, into the list at an index, passed as a second argument

Remove() - removes an object at a specific index, passed as an argument

Clear() - deletes the entire list