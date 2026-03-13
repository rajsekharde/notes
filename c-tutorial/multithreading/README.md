# Multithreading in C

When a single-thread C program is compiled and run, the OS creates a Process. This process starts with a single thread- the Main Thread. All the instructions are executed sequentially.

## multithreading using pthread.h

Library used- pthread.h (POSIX Thread)

Basic implementation:

```bash
// test.c

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

void* action(void* args)
{
    for(int i=1; i<10; i++)
    {
        usleep(300000);
        printf("Thread Count: %d\n", i);
    }
    return NULL;
}

int main()
{
    pthread_t thread;

    pthread_create(&thread, NULL, action, NULL);

    for(int i=1; i<=5; i++)
    {
        usleep(500000);
        printf("Main Count: %d\n", i);
    }

    pthread_join(thread, NULL);

    printf("Thread finished\n");

    return 0;
}

# Output:
rajsekhar@thinkpad:~/notes/c-tutorial/multithreading$ gcc test.c && ./a.out
Thread Count: 1
Main Count: 1
Thread Count: 2
Thread Count: 3
Main Count: 2
Thread Count: 4
Main Count: 3
Thread Count: 5
Thread Count: 6
Main Count: 4
Thread Count: 7
Thread Count: 8
Main Count: 5
Thread Count: 9
Thread finished
rajsekhar@thinkpad:~/notes/c-tutorial/multithreading$ 

```

Code Explanations:

- pthread_t thread : Identifier for a thread

- pthread_create() : Creates new thread and starts executing statements in main & spawned threads immediately. Arguments- pointer to thread, special properties, pointer to void* function, pointer to function argument

- pthread_join() : Pauses main thread execution till spawned thread completes execution. Frees up the thread's resources after ending it

- All statements between pthread_create() and pthread_join() in main function are executed alongside spawned thread

- Thread function must have return type void* and a single void* argument, that can be a pointer to any data type

