import numpy as np
import matplotlib.pyplot as pl

def dft(x, inv=False):
    X = np.empty_like(x, dtype=np.complex128)
    N = len(x)
    for k in range(N):
        S = 0 + 0j
        for n in range(N):
            omega = 2 * np.pi / N * k * n
            if not inv:
                S += (np.cos(omega) - 1j * np.sin(omega)) * x[n]
            else:
                S += (np.cos(omega) + 1j * np.sin(omega)) * x[n]
        if inv:
            S /= N
        X[k] = S
    return X

# Set up matplotlib inline and figure size
pl.rcParams["figure.figsize"] = (15, 7)

# Define parameters
omega = 2 * np.pi
dt = 0.7
t = np.arange(0, 480, dt)
x = np.cos(t * omega / 10 + 1) + np.cos(t * omega / 40 + np.pi / 2)

# Perform DFT
N = len(x)
X = dft(x)
nu = np.arange(N) / (dt * N)
np.seterr(divide='ignore')  # Ignore divide-by-zero warnings
T = 1 / nu

# Plotting
fig = pl.figure()

# Time-domain signal
ax = fig.add_subplot(3, 1, 1)
pl.plot(t, x)
pl.grid()

# Amplitude spectrum
ax = fig.add_subplot(3, 1, 2)
A = np.sqrt(np.real(X) ** 2 + np.imag(X) ** 2) / N
pl.semilogy(T[0:N // 2], A[0:N // 2])
pl.xlim([0, 100])
pl.xticks(np.arange(0, 100, step=10))
pl.grid()

# Phase spectrum
P = np.arctan2(np.imag(X), np.real(X))
ax = fig.add_subplot(3, 1, 3)
pl.plot(T[0:N // 2], P[0:N // 2])
pl.xlim([0, 100])
pl.xticks(np.arange(0, 100, step=10))
pl.grid()

pl.savefig("images/dft_analysis.png")

X_inv = dft(X, inv=True)
fig = pl.figure()
pl.plot(t, np.real(X_inv))
pl.grid()
pl.savefig("images/dft_inverse.png")