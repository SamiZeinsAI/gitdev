const clientId = '88c7fe2b6f54a8507680';
const redirectUri = 'http://localhost:8080/v1/auth/github/callback';
const scope = 'user'; 

export default function handleLogin() {
    
    const authorizationUrl = `https://github.com/login/oauth/authorize?client_id=${clientId}&redirect_uri=${encodeURIComponent(redirectUri)}&scope=${scope}`;
    window.location.href = authorizationUrl;
};