// js search

function performSearch() {
    const searchField = document.getElementById('searchField').value.trim();
    const resultsDiv = document.getElementById('results');

    // Clear previous results
    resultsDiv.innerHTML = '';

    // Check if searchField is empty
    if (searchField === '') {
        resultsDiv.innerHTML = '<p>Search is empty.</p>';
        return;
    }

    console.log(`Performing search for: "${searchField}"`);

    // Send the search query
    fetch(`/api/search/users/?q=${encodeURIComponent(searchField)}`, {
        credentials: 'same-origin'
    })
        .then(response => {
            if (!response.ok) {
                // Handle HTTP errors
                if (response.status === 401 || response.status === 403) {
                    // User is not authenticated, reroute
                    window.location.href = '/login/';
                    throw new Error('User not authenticated');
                } else {
                    resultsDiv.innerHTML = '<p>An error occurred while searching.</p>';
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
            }
            return response.json();
        })
        .then(users => {
            console.log('Users received:', users);
            if (Array.isArray(users) && users.length > 0) {
                users.forEach(user => {
                    // New div for each user
                    const userDiv = document.createElement('div');
                    userDiv.classList.add('profile-item');

                    // Reroute when clicked to profile page
                    userDiv.addEventListener('click', function() {
                        window.location.href = `/profile/${encodeURIComponent(user.username)}`;
                    });


                    // profile pic
                    const profilePicURL = user.profile_pic_url;

                    // html with user info
                    userDiv.innerHTML = `
                        <img src="${profilePicURL}" alt="Profile picture of ${user.displayname}" class="profile-pic">
                        <div class="profile-info">
                            <span class="profile-name">${user.displayname} (@${user.username})</span>
                        </div>
                    `;
                    resultsDiv.appendChild(userDiv);
                });
            } else {
                // if no user found
                const noResultsMessage = document.createElement('p');
                noResultsMessage.textContent = `No user starting with "${searchField}" exists.`;
                resultsDiv.appendChild(noResultsMessage);
            }
        })
        .catch(error => {
            console.error('Error fetching search results:', error);
            if (!resultsDiv.innerHTML) {
                resultsDiv.innerHTML = '<p>An error occurred while searching.</p>';
            }
        });
}

// Event listenre for clicking the search button
document.getElementById('searchButton').addEventListener('click', function() {
    performSearch();
});

// Even listener for enter in search
document.getElementById('searchField').addEventListener('keydown', function(event) {
    if (event.key === 'Enter') {
        performSearch();
    }
});
