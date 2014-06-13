CONFIG = {
  "project": "blink",
  "source_vcs": "svn",
  "source_url": "svn://svn.chromium.org/chrome/trunk/src",
  "status_url": "https://blink-status.appspot.com",
  "masters": {
    "chromium.win": {
      "base_url": "https://build.chromium.org/p/chromium.win",
      "builders": {
        'Win Builder (dbg)',
      },
    },  # chromium.win
    "chromium.mac": {
      "base_url": "https://build.chromium.org/p/chromium.mac",
      "builders": {
        'Mac Builder (dbg)',
      },
    },  # chromium.mac
    "chromium.linux": {
      "base_url": "https://build.chromium.org/p/chromium.linux",
      "builders": {
        'Linux Builder (dbg)',
        'Linux Builder (dbg)(32)',
        'Android Builder (dbg)',
        'Android Builder',
      },
    },  # chromium.linux
    "chromium.webkit": {
      "base_url": "https://build.chromium.org/p/chromium.webkit",
      "builders": {
        'WebKit Win Builder (deps)',
        'WebKit Mac Builder (deps)',
        'WebKit Linux (deps)',
      },
    },  # chromium.webkit
  },
}
