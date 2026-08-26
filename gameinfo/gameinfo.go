package gameinfo

import (
	"wadverifier/gameinfo/doom"
	"wadverifier/gameinfo/doom2"
	"wadverifier/gameinfo/finaldoom"
	"wadverifier/gameinfo/freedoom"
	"wadverifier/gameinfo/heretic"
	"wadverifier/gameinfo/hexen"
	"wadverifier/gameinfo/misc"
	"wadverifier/gameinfo/strife"
	"wadverifier/wad"
)

func PopulateWadInfo() []wad.Entry {

	var list = []wad.Entry{}
	list = append(list, doom.Populate()...)
	list = append(list, doom2.Populate()...)
	list = append(list, finaldoom.Populate()...)
	list = append(list, freedoom.Populate()...)
	list = append(list, heretic.Populate()...)
	list = append(list, hexen.Populate()...)
	list = append(list, strife.Populate()...)
	list = append(list, misc.Populate()...)
	return list

}
