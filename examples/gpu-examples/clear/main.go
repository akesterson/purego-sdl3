package main

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}
	defer sdl.Quit()

	device := sdl.CreateGPUDevice(sdl.GPUShaderFormatSPIRV|sdl.GPUShaderFormatDXIL|sdl.GPUShaderFormatMSL, false, sdl.GetGPUDriver(0))
	defer sdl.DestroyGPUDevice(device)

	window := sdl.CreateWindow("gpu clear screen example", 640, 480, sdl.WindowResizable)
	if window == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)

	if !sdl.ClaimWindowForGPUDevice(device, window) {
		panic(sdl.GetError())
	}
	defer sdl.ReleaseWindowFromGPUDevice(device, window)

Outer:
	for {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				break Outer
			case sdl.EventKeyDown:
				if event.Key().Scancode == sdl.ScancodeEscape {
					break Outer
				}
			}
		}

		commandBuffer := sdl.AcquireGPUCommandBuffer(device)
		if commandBuffer == nil {
			panic(sdl.GetError())
		}

		var swapchainTexture *sdl.GPUTexture
		if !sdl.WaitAndAcquireGPUSwapchainTexture(commandBuffer, window, &swapchainTexture, nil, nil) {
			panic(sdl.GetError())
		}

		if swapchainTexture != nil {
			colorTargetInfo := sdl.GPUColorTargetInfo{}
			colorTargetInfo.Texture = swapchainTexture
			colorTargetInfo.ClearColor = sdl.FColor{0.5, 0.1, 0.3, 1}
			colorTargetInfo.LoadOp = sdl.GPULoadOpClear
			colorTargetInfo.StoreOp = sdl.GPUStoreOpStore

			renderPass := sdl.BeginGPURenderPass(commandBuffer, &colorTargetInfo, 1, nil)
			sdl.EndGPURenderPass(renderPass)
		}

		if !sdl.SubmitGPUCommandBuffer(commandBuffer) {
			panic(sdl.GetError())
		}
	}
}
