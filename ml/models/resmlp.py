import torch
import torch.nn as nn
from typing import Dict, Any
import logging

logger = logging.getLogger(__name__)

class ResMLP(nn.Module):
    def __init__(self, config: Dict[str, Any]):
        super().__init__()
        # TODO: build layers per config
        # example: config = {'input_dim': 30, 'hidden_dims': [128, 64], 'output_dim': 1}
        pass

    def forward(self, x: torch.Tensor) -> torch.Tensor:
        """Forward pass"""
        # TODO: implement
        raise NotImplementedError

    def training_step(self, batch):
        """Single training step for a batch"""
        # TODO: implement
        raise NotImplementedError

    def predict_forward_eps(self, features: torch.Tensor) -> torch.Tensor:
        """Return model prediction for forward EPS"""
        # TODO: implement
        raise NotImplementedError

    def save(self, path: str) -> None:
        """Save model weights"""
        torch.save(self.state_dict(), path)

    @classmethod
    def load(cls, path: str, config: Dict[str, Any]):
        """Load model weights into instance"""
        model = cls(config)
        model.load_state_dict(torch.load(path))
        return model