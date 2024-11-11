# Task Tracker CLI

A simple [CLI tool](https://roadmap.sh/projects/task-tracker) to track all the tasks that you need to, the tasks that you have done and the tasks that you are currently working on.

## How To Use

### build cli

```sh
make build
```

### run cli

#### add new task

```sh
./task-cli add "TASK_NAME"
```

#### update task

```sh
./task-cli update TASK_ID
```

#### delete task

```sh
./task-cli delete TASK_ID
```

#### update task status to in-progress

```sh
./task-cli mark-in-progress TASK_ID
```

#### update task status to done

```sh
./task-cli mark-done TASK_ID
```

#### list all tasks

```sh
./task-cli list
```

#### list tasks marked as todo

```sh
./task-cli list todo
```

#### list tasks marked as in-progress

```sh
./task-cli list in-progress
```

#### list tasks marked as done

```sh
./task-cli list done
```

### run tests

```sh
make test
```

### run lint

```sh
make lint
```
