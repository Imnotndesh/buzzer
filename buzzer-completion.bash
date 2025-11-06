#!/bin/bash

# Bash completion for the 'buzzer' command.

_buzzer_completions() {
    # COMP_CWORD is the index of the current word being completed.
    # COMP_WORDS is an array of the words in the current command line.
    # COMP_REPLY is an array where we put the possible completions.
    local cur_word prev_word
    cur_word="${COMP_WORDS[COMP_CWORD]}"
    prev_word="${COMP_WORDS[COMP_CWORD-1]}"

    # Define all possible subcommands
    local commands="store edit wake get broadcast list remove"

    # Define commands that expect a stored alias as the next argument
    local alias_commands="edit get wake remove"

    # If the previous word is one of the alias commands, we need to complete with an alias.
    if [[ " ${alias_commands} " =~ " ${prev_word} " ]]; then
        # Use our hidden command to get a list of aliases from the database.
        local aliases
        aliases=$(buzzer list-raw 2>/dev/null)
        # Use compgen to filter aliases based on the current word.
        COMPREPLY=($(compgen -W "${aliases}" -- "${cur_word}"))
        return 0
    fi

    # Otherwise, complete with one of the subcommands
    COMPREPLY=($(compgen -W "${commands}" -- "${cur_word}"))
    return 0
}

# Register the completion function for the 'buzzer' command.
complete -F _buzzer_completions buzzer