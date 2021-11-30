# Project CLI
> _This project is still a work in progress but the mvp is working_

Are you tired of forgetting which commands to run or are you annoyed by the to-long-commands-that-do-not-make-any-sense.

Then this is something for you.
 
Project-cli will give you a nice simple to maintain file in which you can document the commands that you need to run.
And the best part is that you can choose the name of the command

## Run Locally

Clone the project

```bash
  git clone https://github.com/koenverburg/project-cli
```

Go to the project directory

```bash
  cd project-cli
```

Install dependencies

```bash
  go install
```

Create a binary

```bash
  go build
```


## Usage/Examples

Create a `.project.yaml` file in your project root. Looking at the example below we could run this by calling `project setup`. This will then run through the commands that listed for the setup section.

```yaml
setup: # This is the name of your command
  - rm -rf node_modules
  - yarn install

```


## License

[MIT](https://choosealicense.com/licenses/mit/)
