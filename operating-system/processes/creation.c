#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>

int main()
{
    printf("Parent process (PID: %d) is starting\n", getpid());

    pid_t pid = fork();

    if(pid < 0)
    {
        // Fork failed
        fprintf(stderr, "Fork Failed\n");
        return 1;
    }
    else if(pid == 0)
    {
        // Child process path
        printf("CHILD: PID: %d, Parent PID: %d\n", getpid(), getppid());
        printf("CHILD: Sleeping for 20 seconds\n");
        sleep(20);
        printf("CHILD: Terminating now\n");
        exit(0);
    }
    else
    {
        /// Parent process path
        printf("PARENT: PID: %d, Child PID: %d\n", getpid(), pid);
        printf("PARENT: Waiting for child process to finish\n");
        wait(NULL);
        printf("PARENT: Child process terminated. Parent process exiting\n");
    }

    return 0;
}