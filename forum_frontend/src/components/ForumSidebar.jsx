import React from 'react';
import './ForumSidebar.css';

export default function ForumSidebar() {
  return (
    <aside className="forum-sidebar">
      <button className="start-discussion">Start a Discussion</button>
      <div className="sidebar-section">
        <button className="latest-btn">Latest ▼</button>
      </div>
      <nav className="sidebar-nav">
        <ul>
          <li className="active"><span role="img" aria-label="discussions">💬</span> All Discussions</li>
          <li><span role="img" aria-label="tags">🔲</span> Tags</li>
          <li><span role="img" aria-label="general">⬛</span> General</li>
        </ul>
      </nav>
    </aside>
  );
} 