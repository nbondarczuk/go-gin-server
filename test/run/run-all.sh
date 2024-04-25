#!/bin/bash
export PATH=.:$PATH
read-health.sh
echo
drop-entity-tag-collection.sh
echo
create-entity-tag.sh
echo
read-entity-tag.sh
echo
drop-entity-tag-collection.sh
echo

