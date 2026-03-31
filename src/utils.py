import logging
import os
from typing import Dict, Any

def get_config() -> Dict[str, Any]:
    config = {
        'REDIS_HOST': os.environ.get('REDIS_HOST', 'localhost'),
        'REDIS_PORT': os.environ.get('REDIS_PORT', '6379'),
        'REDIS_DB': os.environ.get('REDIS_DB', 0),
        'CACHE_TTL': int(os.environ.get('CACHE_TTL', 60)),
        'LOG_LEVEL': os.environ.get('LOG_LEVEL', 'INFO')
    }
    return config

def configure_logging(log_level: str) -> None:
    logging.basicConfig(level=getattr(logging, log_level.upper()))