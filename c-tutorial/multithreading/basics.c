#include <stdio.h>
#include <pthread.h>
#include <unistd.h>

typedef struct
{
    int start;
    int end;
} thread_data;

// Thread function must return void* and have single void* argument
void* action(void* args)
{
    thread_data td = *(thread_data*)args; // Convert argument to appropriate data type
    for(int i=td.start; i<=td.end; i++)
    {
        usleep(300000);
        printf("Count: %d\n", i);
    }
}

int main()
{
    // Data is passed to thread function as a struct
    thread_data td1 = {1, 10};

    // Identifier to thread
    pthread_t thread;

    // Create new thread and execute specified function
    pthread_create(&thread, NULL, action, &td1);

    // Execute main thread statements simultaenously
    printf("Main thread\n");
    for(int i=1; i<=5; i++)
    {
        usleep(500000);
        printf("Main thread count: %d\n", i);
    }

    // Pause main thread till spawned thread completes execution & free its resources
    pthread_join(thread, NULL);

    printf("Thread finished\n");
    return 0;
}

/*
pthread_create(&thread_id, special-properties, function, pointer-to-argument)
special-properties - NULL for default properties
argument - NULL if no arguments are passed
*/