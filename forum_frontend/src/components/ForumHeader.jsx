import React from 'react';
import './ForumHeader.css';

export default function ForumHeader() {
  return (
    <header className="forum-header">
      <div className="forum-header-left">
        <span className="forum-logo">Flarum/demo</span>
      </div>
      <div className="forum-header-center">
        <input className="forum-search" type="text" placeholder="Search Forum" />
      </div>
      <div className="forum-header-right">
        <button className="forum-signup">Sign Up</button>
        <button className="forum-login">Log In</button>
      </div>
    </header>
  );
} 