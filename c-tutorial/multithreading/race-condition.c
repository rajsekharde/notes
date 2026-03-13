#include <stdio.h>
#include <pthread.h>

/*
When multiple threads try to access or modify a shared resource simultaenously, race condition occurs.
Example: on calling normal_run(), expected output should be val = 10*10000 = 10000. But it's always
less than 100k. This occurs because multiple threads may try to read and increment the same value of
val at the same time, causing only a single increment to take place.

Solution: use mutex (mutual exclusion) or a lock that allows only one thread to touch the shared
resources at a time. When a thread locks the resource, other threads wait till it's unlocked.
The code between lock and unlock should be as short as possible.
*/

int val = 0;
int val2 = 0;
pthread_mutex_t lock;

void* action_normal(void* args)
{
    for(int i=0; i<10000; i++)
    {
        val += 1;
    }
}

void *action_mutex(void *args)
{
    for (int i = 0; i < 10000; i++)
    {
        pthread_mutex_lock(&lock); // lock the resource
        val2 += 1;
        pthread_mutex_unlock(&lock); // unlock the resource
    }
}

int normal_run()
{
    int tc = 10;
    pthread_t threads[tc];

    for(int i=0; i<tc; i++)
    {
        pthread_create(&threads[i], NULL, action_normal, NULL);
    }

    for(int i=0; i<tc; i++)
    {
        pthread_join(threads[i], NULL);
    }

    printf("val = %d\n", val);

    return 1;
}

int mutex_run()
{
    int tc = 10;
    pthread_t threads[tc];

    pthread_mutex_init(&lock, NULL); // initialize mutex

    for (int i = 0; i < tc; i++)
    {
        pthread_create(&threads[i], NULL, action_mutex, NULL);
    }

    for (int i = 0; i < tc; i++)
    {
        pthread_join(threads[i], NULL);
    }

    pthread_mutex_destroy(&lock); // destroy mutex

    printf("val = %d\n", val2);

    return 1;
}

int main()
{
    normal_run();

    mutex_run();

    return 0;
}