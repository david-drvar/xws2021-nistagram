import React, {useEffect, useState} from 'react';
import ProfileForSug from "../HomePage/ProfileForSug";
import ProfileForAutocomplete from "./ProfileForAutocomplete";


function UserAutocomplete(props) {
    const [activeSuggestion, setActiveSuggestion] = useState(0);
    const [filteredSuggestions, setFilteredSuggestions] = useState([]);
    const [showSuggestions, setShowSuggestions] = useState(false);
    const [userInput, setUserInput] = useState("");

    function onClick(suggestion) {
        props.setSelectedUser && props.setSelectedUser(suggestion);
        props.setSelectedUser && setUserInput(suggestion.username);
        props.setSelectedUser && setShowSuggestions(false);

        props.addToTaglist &&setActiveSuggestion(0);
        props.addToTaglist && setShowSuggestions(false);
        props.addToTaglist && setUserInput("");
        props.addToTaglist && setFilteredSuggestions([]);
        props.addToTaglist && props.addToTaglist(suggestion);
    }

    function onChange(e) {
        const { suggestions } = props;
        const userInput = e.currentTarget.value;
        //console.log(suggestions);
        const filteredSuggestions = suggestions.filter(
            (suggestion) =>
                suggestion.username.toLowerCase().indexOf(userInput.toLowerCase()) > -1
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
                                <ProfileForAutocomplete image={suggestion.profilePhoto} username={suggestion.username} firstName={suggestion.firstName} lastName={suggestion.lastName}  caption={suggestion.biography} urlText="Follow" iconSize="medium" captionSize="small" storyBorder={true} />
                                {/*{suggestion}*/}
                            </li>
                        );
                    })}
                </ul>
            );
        } else {
            suggestionsListComponent = (
                <div class="no-suggestions">
                    <em>No suggestions!</em>
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

}export default UserAutocomplete;