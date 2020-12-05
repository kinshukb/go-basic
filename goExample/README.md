Hi there,
It was exactly around 2.6 years back in June 2018, when I joined my first job as a Software Engineer, I was introduced to GO.
After a day or two of reading docs, learning the languages' basics I jumped into implementing some coding examples to get a better understading.

This might help programmers at beginner level, I've given some comparisions to other languages and how those are achieved in go, there are examples with comments to help understand the concepts and programming paradigm - the go way.
Although these can't be compared to idiomatic go, which is the ideal way to code in go but it can help at crucial points where newbies stuck, read comments!

First Impressions of Go:
Go is a fast, object oriented language (It has some features you will often see in object oriented programming languages, and other features you will often see in functional programming languages, but I think ultimately it counts as a procedural language. It's basically along the same lines as C, but a bit more higher level and with some more modern language concepts built in from the beginning.)
It is built on top of c,
Features like concurrency, polymorphism, garbage collection are supported internally by the go-runtime.
no intermediary code like java or .net
that's why the speed (also pointers exist)

no concept of classes instead structs and funcs are used to conceptualize a class 
func() with (ref/var structType ) 

polymorphism is achieved by using structs and interfaces together

a function with argument as interface type is used to demonstrate polymorphic nature

REFER subStructs.go --> recommended to understand structs and interfaces

check CompositionInheri.go to see how inheritance is achieved in golang there is no extending of classes 
instead inheritance can be achieved using composition only.

if-else and switch has a particular syntax to be followed
check it out,
package concept is similar to C++'s namespace and javas' pkg.


go-routines are separate path of execution just like a thread but are lightweight as these are not actually OS threads 
In fact there is a thread pool of goroutines and a goroutine runs on top of OS thread
channels acts as a pipe enabling communication between two different threads
internally these are just mutexes with a lock and data.
The go-runtime scheduler is responsible for managing them.


Go code must be kept inside a workspace. A workspace is a directory hierarchy with three directories at its root:

src contains Go source files organized into packages (one package per directory),

pkg contains package objects, and

bin contains executable commands.