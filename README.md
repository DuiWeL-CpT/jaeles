
<p align="center">
  <img alt="Jaeles" src="https://image.flaticon.com/icons/svg/1432/1432425.svg" height="140" />
  <p align="center">
    <a href=""><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
    <a href="https://github.com/jaeles-project/jaeles"><img alt="Release" src="https://img.shields.io/badge/version-beta%20v0.1-red.svg"></a>
  </p>
</p>

**Jaeles** is a powerful, flexible and easily extensible framework written in Go for building your own Web Application Scanner.

![Architecture](https://github.com/jaeles-project/jaeles-plugins/blob/master/imgs/jaeles-architecture.png?raw=true)

## Installation

```
go get -u github.com/jaeles-project/jaeles
```

Please visit the [Official Documention](https://jaeles-project.github.io/) for more details.

Checkout [Signature Repo](https://github.com/jaeles-project/jaeles-signatures) for base signature.

## Usage
More usage [here](https://jaeles-project.github.io/usage/)

Example commands.
```
jaeles scan -u http://example.com

jaeles scan -s signatures/common/phpdebug.yaml -U /tmp/list_of_urls.txt

jaeles scan --retry 3 --verbose -s "signatures/cves/jira-*" -U /tmp/list_of_urls.txt

jaeles --verbose server -s sqli
```

## Showcases
More showcase [here](https://jaeles-project.github.io/showcases/)

[![asciicast](https://asciinema.org/a/281205.svg)](https://asciinema.org/a/281205)
<p align="center">
Detect Jira SSRF CVE-2019-8451
</p>

## Contribute

If you have some new idea about this project, issue, feedback or found some valuable tool feel free to open an issue for just DM me via @j3ssiejjj.

## License

`Jaeles` is made with ♥  by [@j3ssiejjj](https://twitter.com/j3ssiejjj) and it is released under the MIT license.