name: vue

on:
push:
branches:
- main
- books-ui
pull_request:
branches:
- main
- books-ui


jobs:
lint-books-ui:
runs-on: ubuntu-latest

    steps:
      - name: log
        uses: echo github.ref

      - name: Extract branch name
        run: echo "BRANCH=${GITHUB_REF#refs/heads/}" >> $GITHUB_ENV

      - name: Checkout
        uses: actions/checkout@v4.1.1

      - name: cd books-ui
        run: cd books-ui

      - name: lint
      - uses: stefanoeb/eslint-action@1.0.2
        with:
          files: src/


test-books-ui:
runs-on: ubuntu-latest

    permissions:
      contents: read

    steps:
      - Checkout
      - uses: actions/checkout@v4

      - cd books-ui
      - run: cd books-ui

      - name: 'Install Node'
        uses: actions/setup-node@v4
        with:
          node-version: '20.x'

      - name: 'Install Deps'
        run: npm install

      - name: 'Test'
        run: npx vitest
