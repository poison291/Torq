## 🚧 Under Development

The following features are currently being developed or planned for upcoming versions:

### Core Features (Backend)
- Torrent metadata generation for single and multi-file torrents
- Magnet link generation and parsing
- Peer-to-peer connection handling and seeding logic
- Tracker support + DHT for peer discovery
- Real-time stats (upload/download speeds, peers, pieces)

### Frontend (UI)
- Modern desktop UI using Wails + React
- File picker and drag-and-drop support
- Torrent list with progress bars, peer count, and status
- Action buttons: Start/Stop Seeding, Copy Magnet Link
- Collapsible activity/log panel with live updates
- Light/Dark theme support

### Integration
- Backend API exposed to frontend via Wails bridge
- Error handling and user notifications
- Cross-platform build support (Windows, macOS, Linux)

### Optional / Future Enhancements
- Multiple torrent management (batch upload/download)
- Custom piece size selection
- Advanced tracker options
- Auto-update mechanism for the desktop app
- Optional cloud seeding / analytics features for premium users
