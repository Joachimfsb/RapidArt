// js frontpage

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
            // Toggle Like State
            button.dataset.liked = isLiked ? '0' : '1';
            button.querySelector('img').src = newIcon;

            // Update like count
            const likeCount = button.querySelector('.like-count');
            likeCount.textContent = parseInt(likeCount.textContent) + (isLiked ? -1 : 1);
        }
    };

    xhr.open('POST', endpoint, true);
    xhr.send();
}
