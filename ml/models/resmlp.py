import torch
import torch.nn as nn


class ResMLP(nn.Module):
    def __init__(self, config):
        super().__init__()
        # TODO: build layers per config
        pass

    def forward(self, x):
        """Forward pass"""
        # TODO: implement
        raise NotImplementedError

    def training_step(self, batch):
        """Single training step for a batch"""
        # TODO: implement
        raise NotImplementedError

    def predict_forward_eps(self, features):
        """Return model prediction for forward EPS"""
        # TODO: implement
        raise NotImplementedError

    def save(self, path):
        """Save model weights"""
        torch.save(self.state_dict(), path)

    @classmethod
    def load(cls, path, config):
        """Load model weights into instance"""
        model = cls(config)
        model.load_state_dict(torch.load(path))
        return model