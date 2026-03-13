#include <stdio.h>
#include <pthread.h>
#include <unistd.h>

/*
Multiple threads can ba spawned by creating and storing their identifiers
in a pthread_t array, then iterating through the array once for creation
and once again for joining
*/

typedef struct
{
    int id;
    int sleep_time;
} thread_data;

void* thread_function(void* args)
{
    // int id = *(int*)args;
    thread_data td = *(thread_data*)args;

    for(int i=1; i<=5; i++)
    {
        // usleep(id*300000);
        // printf("Thread: %d Count: %d\n", id, i);

        usleep(td.id * td.sleep_time);
        printf("Thread: %d Count: %d\n", td.id, i);
    }
}

int main()
{
    int thread_count = 5;
    int sleep_time = 1000000/thread_count;

    // Store all thread identifiers in an array
    pthread_t threads[thread_count];

    int thread_ids[thread_count];
    thread_data tds[thread_count];

    for(int i=0; i<thread_count; i++)
    {
        thread_ids[i] = i+1;
        tds[i] = (thread_data){i+1, sleep_time};

        // Create threads
        // pthread_create(&threads[i], NULL, thread_function, &thread_ids[i]);
        pthread_create(&threads[i], NULL, thread_function, &tds[i]);
    }

    for(int i=0; i<thread_count; i++)
    {
        // Join threads
        pthread_join(threads[i], NULL);
    }

    return 0;
}