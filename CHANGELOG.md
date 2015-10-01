## 0.1.2 (unreleased)

IMPROVEMENTS:

 * app/ruby: Use --no-document when installing bundler in dev [GH-130]
 * app/ruby: Install apt deps based on detected gems [GH-137]
 * app/ruby: Bundle automatically [GH-156]
 * app/php: support customizing PHP version [GH-105]
 * app: support Vagrant parallels provider in dev [GH-85]

BUG FIXES:

 * app/node: fix node download directory in build [GH-125]
 * app/ruby: allow `gem install` to work as `vagrant` user in dev [GH-129]
 * app: fix Vagrant warning about box name [GH-110]

## 0.1.1 (September 28, 2015)

BUG FIXES:

* core: Fix marshalling format problem preventing successful parsing of an Appfile with multiple dependencies declared [GH-83]
* app: Fix string escaping issues affecting `otto dev` on Windows hosts [GH-79]
* app/node: Fix issue preventing `node` apps from interpolating configs properly [GH-73]

## 0.1.0 (September 28, 2015)

* Initial release
