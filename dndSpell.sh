#!/bin/bash


cd /home/mike/scripts/go_projects/dndSpell/spells/
f=$(fzf -i)

dndSpellParser $f
