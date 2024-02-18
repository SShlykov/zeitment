const routes = [
  {"path": "/socket.io",    "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/socket.io/*",  "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/api",          "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/api/*",        "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/api/*/*",      "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/api/*/*/*",    "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/api/*/*/*/*",  "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/swagger",      "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/swagger/*",    "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/swagger/*/*",  "backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/swagger/*/*/*","backend": "http://backend:8000",     "role":  "*", "site": "cdrp"},
  {"path": "/",             "backend": "http://frontend:8000",    "role":  "*", "site": "cdrp"},
  {"path": "/*",            "backend": "http://frontend:8000",    "role":  "*", "site": "cdrp"},
  {"path": "/*/*",          "backend": "http://frontend:8000",    "role":  "*", "site": "cdrp"},
  {"path": "/*/*/*",        "backend": "http://frontend:8000",    "role":  "*", "site": "cdrp"},
  {"path": "/*/*/*/*",      "backend": "http://frontend:8000",    "role":  "*", "site": "cdrp"},
  {"path": "/*/*/*/*/*",    "backend": "http://frontend:8000",    "role":  "*", "site": "cdrp"}
]

const sites = [
  {"name": "cdrp"},
  {"name": "live_monitore"},
  {"name": "buk"},
  {"name": "mtso"},
  {"name": "as_dmps"}
]

export {routes, sites}
