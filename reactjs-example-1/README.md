React.js Example 1
==================

## Setup
### Installation for Node.js

Install [anyenv](https://github.com/riywo/anyenv):

```bash
git clone https://github.com/riywo/anyenv ~/.anyenv
echo 'export PATH="$HOME/.anyenv/bin:$PATH"' >> ~/.bash_profile
echo 'eval "$(anyenv init -)"' >> ~/.bash_profile
exec $SHELL -l
```

Install [ndenv](https://github.com/riywo/ndenv):

```bash
anyenv install ndenv
exec $SHELL -l
```

Install Node.js:

```bash
ndenv install v4.1.1
ndenv global v4.1.1
```

Install gulp:

```bash
npm install -g gulp
ndenv rehash
```

### Cloning this repos

```bash
git clone https://github.com/nownabe/examples.git
cd examples/reactjs-example-1
npm install
```

## Run
Run gulp:

```bash
gulp
```
