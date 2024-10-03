//MIDLERTIDIG script for å illustrere/teste søke funksjon. Endres ved databsasetilkobling

document.getElementById('searchButton').addEventListener('click', function() {
    const searchField = document.getElementById('searchField').value;
    const resultsDiv = document.getElementById('results');
    
    // Fjern forrige resultater
    resultsDiv.innerHTML = '';

    // Mock profiler
    const mockProfiles = [
        { name: 'Alice Johnson', posts: 12 },
        { name: 'Bob Smith', posts: 8 },
        { name: 'Charlie Brown', posts: 5 },
        { name: 'David Wilson', posts: 15 },
        { name: 'Ella Testing', posts: 10 }
    ];

    // Filter (inkluder bokstaver..)
    const filteredProfiles = mockProfiles.filter(profile => 
        profile.name.toLowerCase().includes(searchField.toLowerCase())
    );

    // Vis profiler som matcher filteret, hvis ingen finnes kommer melding om det
    if (filteredProfiles.length > 0) {
        filteredProfiles.forEach(profile => {
            //Oprett en ny div for hver profil
            const profileDiv = document.createElement('div');
            profileDiv.classList.add('profile-item');

            //Setter inn html med profilinfo
            profileDiv.innerHTML = `
                <img src="/res/img/profileblack.png" alt="Profile" class="profile-pic">
                <div class="profile-info">
                    <span class="profile-name">${profile.name}</span>
                    <span class="post-count">Number of posts: ${profile.posts}</span>
                </div>
            `;
            resultsDiv.appendChild(profileDiv); // Legger til infoen i resultatområdet
        });
    } else {
        resultsDiv.textContent = 'No profiles found.';
    }
});
