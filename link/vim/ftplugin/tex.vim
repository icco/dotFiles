" Wrap lines, and break only on whitespace.
setlocal textwidth=0
setlocal wrap
setlocal linebreak

" Use the screen-based versions of the most common movement commands instead of
" the text-based versions.
noremap <buffer> <Down> gj
noremap <buffer> j gj
noremap <buffer> <Up> gk
noremap <buffer> k gk
noremap <buffer> 0 g0
noremap <buffer> ^ g^
noremap <buffer> $ g$
noremap <buffer> G G$
