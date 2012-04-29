function! s:ToDoList ()
    cclose
    let task_list = []
    for row in split(system('ack --column "(TODO|CHANGED|FIXME)"'), '\n')
        let t = split(row, ':')
        let task_dict = {'filename': t[0], 'lnum': t[1], 'col': t[2]}
        let task_dict.text = substitute(join(t[3:-1]), '\s\+', ' ', '')
        let task_list += [task_dict]
    endfor
    call setqflist(task_list, 'r')
    copen
endfunction
command! ToDoList call <SID>ToDoList()
map <silent> <Leader>td :ToDoList<CR>
