# `nac`: nac ain't configuration 

```bash
$ nac wtf /usr/local/

-> Tertiary hierarchy for local data, specific to this host. Typically has further subdirectories, e.g., bin/, lib/, share/.[30]


$ nac wtf /box/www

-> (no standard docs)

On Oct 28, 2009, Andrew Gaffney:
"This directory holds all information pertinent to the Box web application"

On Dec 11, 2012, Andrew Gaffney:
"Nothing else should go in this folder, it is deprecated"


$ nac wtf /box/www/devtools

-> (no standard docs)

On Oct 28, 2009, Anthony Bishopric:
"This is a puppet-managed clone of devtools, do not edit yourself"

Tags: "eng-services", "pushgerrit", "tools"


$ nac wtf -R /

1. /box/www
-> (no standard docs)

On Oct 28, 2009, Andrew Gaffney:
"This directory holds all information pertinent to the Box web application"

On Dec 11, 2012, Andrew Gaffney:
"Nothing else should go in this folder, it is deprecated"

2. /box/www/devtools
-> (no standard docs)

On Oct 28, 2009, Anthony Bishopric:
"This is a puppet-managed clone of devtools, do not edit yourself"

Tags: "eng-services", "pushgerrit", "tools"


$ nac wtf -t "pushgerrit" -R /

-> (no standard docs)

On Oct 28, 2009, Anthony Bishopric:
"This is a puppet-managed clone of devtools, do not edit yourself"

Tags: "eng-services", "pushgerrit", "tools"


$ nac tags

eng-services		4
tools				3
pushgerrit			1


$ nac change "Add Foo extension to PHP"

$(Add Foo extension to PHP) echo "extension=foo.so" > /etc/php.d/ext-foo.ini
$(Add Foo extension to PHP) exit


$ nac search "application"

1. /box/www (directory)
On Oct 28, 2009, Andrew Gaffney:
"This directory holds all information pertinent to the Box web _application_"

2. /box/www/current (directory)
On Oct 28, 2009, Andrew Gaffney:
"This is the conventional Box web _application_ checkout."

```
