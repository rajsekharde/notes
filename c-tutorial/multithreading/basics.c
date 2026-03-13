#include <stdio.h>
#include <pthread.h>
#include <unistd.h>

typedef struct
{
    int start;
    int end;
} thread_data;

void* action(thread_data* td)
{
    for(int i=td->start; i<=td->end; i++)
    {
        usleep(300000);
        printf("Count: %d\n", i);
    }
}

int main()
{
    thread_data td1 = {1, 10};

    pthread_t thread_id;
    pthread_create(&thread_id, NULL, action, &td1);

    printf("Main thread\n");
    for(int i=1; i<=5; i++)
    {
        usleep(500000);
        printf("Main thread count: %d\n", i);
    }

    pthread_join(thread_id, NULL);

    printf("Thread finished\n");
    return 0;
}