""
" @section Introduction, intro
" Core settings for the soy filetype.

if exists('b:did_ftplugin')
  finish
endif
let b:did_ftplugin = 1

let b:undo_ftplugin = 'setlocal comments< formatoptions<'

setlocal formatoptions-=t

setlocal comments=sO:*\ -,mO:*\ \ ,exO:*/,s1:/*,mb:*,ex:*/,://
