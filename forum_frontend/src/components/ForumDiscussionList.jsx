import React from 'react';
import './ForumDiscussionList.css';

const discussions = [
  {
    id: 1,
    avatar: 'V',
    avatarColor: '#b39ddb',
    title: 'Et est ea.',
    user: 'allie_boyer',
    time: '2 days ago',
    replies: 2300,
  },
  {
    id: 2,
    avatar: 'F',
    avatarColor: '#ffe082',
    title: 'Nemo dignissimos sit temporibus nihil recusandae nihil vel aut.',
    user: 'ally28',
    time: '2 days ago',
    replies: 11,
  },
  {
    id: 3,
    avatar: 'N',
    avatarColor: '#ce93d8',
    title: 'Et omnis.',
    user: 'victoria09',
    time: '2 days ago',
    replies: 17,
  },
  {
    id: 4,
    avatar: 'W',
    avatarColor: '#fff59d',
    title: 'Quasi ad sit rerum quis velit.',
    user: 'hassan_watsica',
    time: '2 days ago',
    replies: 12,
  },
  {
    id: 5,
    avatar: 'H',
    avatarColor: '#ffccbc',
    title: 'Reprehenderit molestiae qui.',
    user: 'rosalee_tromp',
    time: '2 days ago',
    replies: 10,
  },
  {
    id: 6,
    avatar: 'W',
    avatarColor: '#80deea',
    title: 'Recusandae sint magni deleniti inventore dolore qui non.',
    user: 'ally28',
    time: '2 days ago',
    replies: 14,
  },
  {
    id: 7,
    avatar: 'R',
    avatarColor: '#f8bbd0',
    title: 'Fugiat sed quis dolorem maxime.',
    user: 'shannon_rempel',
    time: '2 days ago',
    replies: 17,
  },
];

function formatReplies(num) {
  if (num > 999) return (num / 1000).toFixed(1) + 'K';
  return num;
}

export default function ForumDiscussionList() {
  return (
    <main className="forum-discussion-list">
      <button className="refresh-btn" title="Refresh">🔄</button>
      <ul>
        {discussions.map(d => (
          <li key={d.id} className="discussion-item">
            <div className="discussion-avatar" style={{backgroundColor: d.avatarColor}}>{d.avatar}</div>
            <div className="discussion-content">
              <div className="discussion-title">{d.title}</div>
              <div className="discussion-meta">↪ {d.user} replied {d.time}</div>
            </div>
            <div className="discussion-replies">
              <span role="img" aria-label="replies">💬</span> {formatReplies(d.replies)}
            </div>
          </li>
        ))}
      </ul>
    </main>
  );
} 