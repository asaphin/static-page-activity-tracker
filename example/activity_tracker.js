function sendActivity(page, data) {
    const apiUrl = 'https://5lldnv5x76.execute-api.us-east-1.amazonaws.com/prod/activity';
    const requestBody = {
        page: page,
        data: data
    };

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'x-api-key': 'jTqwugBBxpa9XRpIxIyZu1naFsWvD2YS88T8gG8S',
        },
        body: JSON.stringify(requestBody),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log('Activity sent successfully:', data);
        })
        .catch(error => {
            console.error('Error sending activity:', error);
        });
}

document.addEventListener('DOMContentLoaded', function() {
    const pageName = window.location.hostname;
    const eventData = {
        pageURL: window.location.href,
        pageTitle: document.title,
        userAgent: navigator.userAgent,
    };

    sendActivity(pageName, eventData);
});
