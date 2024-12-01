document.addEventListener('DOMContentLoaded', function () {
    // Feed toggle buttons
    const toggleFollowed = document.querySelector('#toggle-followed');
    const toggleGlobal = document.querySelector('#toggle-global');

    // Initialize active state
    const activeFeed = window.location.href.includes('feed=global') ? 'global' : 'followed';
    if (activeFeed === 'global') {
        toggleGlobal.classList.add('active');
        toggleFollowed.classList.add('inactive');
    } else {
        toggleFollowed.classList.add('active');
        toggleGlobal.classList.add('inactive');
    }

    toggleFollowed.addEventListener('click', () => switchFeed('followed', toggleFollowed, toggleGlobal));
    toggleGlobal.addEventListener('click', () => switchFeed('global', toggleGlobal, toggleFollowed));

    // Like buttons
    const likeButtons = document.querySelectorAll('.like-wrapper');
    likeButtons.forEach((button) => button.addEventListener('click', toggleLike));

    // Report buttons
    const reportButtons = document.querySelectorAll('.report-wrapper');
    reportButtons.forEach((button) => button.addEventListener('click', handleReport));
});

// Switch Feed
function switchFeed(feedType, activate, deactivate) {
    activate.classList.add('active');
    activate.classList.remove('inactive');
    deactivate.classList.remove('active');
    deactivate.classList.add('inactive');

    window.location.href = `/?feed=${feedType}`;
}

// Toggle Like
function toggleLike(event) {
    const button = event.currentTarget;
    const postId = button.dataset.postId;
    const isLiked = button.dataset.liked === '1';

    const xhr = new XMLHttpRequest();
    const endpoint = isLiked ? `/api/post/unlike/${postId}` : `/api/post/like/${postId}`;
    const newIcon = isLiked ? '/res/icon/heart.svg' : '/res/icon/heart-fill-red.svg';

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 204) {
            button.dataset.liked = isLiked ? '0' : '1';
            button.querySelector('img').src = newIcon;

            const likeCount = button.querySelector('.like-count');
            likeCount.textContent = parseInt(likeCount.textContent) + (isLiked ? -1 : 1);
        }
    };

    xhr.open('POST', endpoint, true);
    xhr.send();
}

// Handle Report
function handleReport(event) {
    const button = event.currentTarget;
    const postId = button.dataset.postId;
    const isReported = button.dataset.reported === '1'; // Check if already reported

    if (isReported) {
        // Provide feedback if the post is already reported
        alert("You have already reported this post.");
        return;
    }

    // Prompt for report reason
    let message = "";
    while (!message.trim()) {
        message = prompt("Enter your reason for reporting this post (required):");
        if (!message) {
            alert("A reason is required to submit a report.");
        }
    }

    // Send report request
    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 204) {
                alert("Report submitted successfully.");

                // Change svg once reported
                button.dataset.reported = '1';
                button.querySelector('img').src = '/res/icon/flag-fill.svg';
            } else {
                alert("Failed to submit the report. Please try again.");
            }
        }
    };

    xhr.open("POST", `/api/post/report/${postId}`, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify({ message }));
}

