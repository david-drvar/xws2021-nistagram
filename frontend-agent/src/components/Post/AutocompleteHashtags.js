import React, {useEffect, useState} from 'react';


function AutocompleteHashtags(props) {
    const [activeSuggestion, setActiveSuggestion] = useState(0);
    const [filteredSuggestions, setFilteredSuggestions] = useState([]);
    const [showSuggestions, setShowSuggestions] = useState(false);
    const [userInput, setUserInput] = useState("");

    function onClick(suggestion) {
        setActiveSuggestion(0);
        setShowSuggestions(false);
        setUserInput("");
        setFilteredSuggestions([]);

        props.addToHashtaglist(suggestion);
    }

    function onClickNewSuggestion() {
        setActiveSuggestion(0);
        setShowSuggestions(false);
        setUserInput("");
        setFilteredSuggestions([]);

        props.handleHashtagAutocompleteNewSuggestion(userInput);
    }

    function onChange(e) {
        const { suggestions } = props;
        const userInput = e.currentTarget.value;
        //console.log(suggestions);
        const filteredSuggestions = suggestions.filter(
            (suggestion) =>
                suggestion.text.toLowerCase().indexOf(userInput.toLowerCase()) > -1
        );
        //console.log(e);
        setActiveSuggestion(0);
        setShowSuggestions(true);
        setUserInput(e.currentTarget.value)
        setFilteredSuggestions(filteredSuggestions);
    }

    function onKeyDown(e) {
        if (e.keyCode === 13) {
            setActiveSuggestion(0);
            setShowSuggestions(false);
            setUserInput(filteredSuggestions[activeSuggestion])
        }
        else if (e.keyCode === 38) {
            if (activeSuggestion === 0) {
                return;
            }
            setActiveSuggestion(activeSuggestion - 1);
        }
        else if (e.keyCode === 40) {
            if (activeSuggestion - 1 === filteredSuggestions.length) {
                return;
            }
            setActiveSuggestion(activeSuggestion + 1);
        }
    }

    let suggestionsListComponent;
    if (showSuggestions && userInput) {
        if (filteredSuggestions.length) {
            suggestionsListComponent = (
                <ul class="suggestions">
                    {filteredSuggestions.map((suggestion, index) => {
                        return (
                            <li  key={suggestion} onClick={() => onClick(suggestion)}>
                                {suggestion.text}
                            </li>
                        );
                    })}
                </ul>
            );
        } else {
            suggestionsListComponent = (
                <div class="no-suggestions" onClick={() => onClickNewSuggestion()}>
                    <em>No suggestions!</em><br/>
                    Add your custom tag - #{userInput}
                </div>
            );
        }
    }

    return (
        <React.Fragment>
            <input
                type="text"
                onChange={onChange}
                onKeyDown={onKeyDown}
                value={userInput}
            />
            {suggestionsListComponent}
        </React.Fragment>
    );

}export default AutocompleteHashtags;