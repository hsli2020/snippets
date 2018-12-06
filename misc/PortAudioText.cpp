// PortAudioTest.cpp : 定义控制台应用程序的入口点。
// https://blog.csdn.net/xzpblog/article/details/79481420

#include "stdafx.h"
#include <iostream>
#include "portAudio/portaudio.h"

using namespace std;

#pragma comment(lib,"portAudio/portaudio_x86.lib")

PaStream *in_stream;
PaStream *out_stream;

int paOutStreamBk(const void* input, void* output, unsigned long frameCount,
    const PaStreamCallbackTimeInfo* timeInfo, PaStreamCallbackFlags statusFlags, void * userData)
{
    Pa_WriteStream(out_stream, input, frameCount);
    return 0;
}

bool StartPlay()
{
    PaStreamParameters ouputParameters;
    PaError err = paNoError;
    ouputParameters.device = Pa_GetDefaultOutputDevice();
    ouputParameters.channelCount = 2;
    ouputParameters.sampleFormat = paFloat32;
    ouputParameters.suggestedLatency = Pa_GetDeviceInfo(ouputParameters.device)->defaultLowOutputLatency;
    ouputParameters.hostApiSpecificStreamInfo = NULL;

    err = Pa_OpenStream(&out_stream, NULL, &ouputParameters, 44100, 256, paFramesPerBufferUnspecified, NULL, NULL);
    if (err != paNoError) {
        return false;
    }

    err = Pa_StartStream(out_stream);
    if (err != paNoError) {
        return false;
    }
    return true;
}

bool InitPortAudio(int deviceIndex)
{
    PaStreamParameters intputParameters;
    PaError err = paNoError;

    err = Pa_Initialize();
    if (err != paNoError) goto error;

    if (deviceIndex < 0)
    {
        deviceIndex = Pa_GetDefaultInputDevice();
    }
    intputParameters.device = deviceIndex;
    intputParameters.channelCount = 2;
    intputParameters.sampleFormat = paFloat32;
    intputParameters.suggestedLatency = Pa_GetDeviceInfo(intputParameters.device)->defaultLowInputLatency;
    intputParameters.hostApiSpecificStreamInfo = NULL;

    err = Pa_OpenStream(&in_stream, &intputParameters, NULL, 44100, 256, paFramesPerBufferUnspecified, paOutStreamBk, NULL);
    if (err != paNoError) {
        return false;
    }
    return StartPlay();

error:
    Pa_Terminate();
    fprintf(stderr, "An error occured while using the portaudio stream\n");
    fprintf(stderr, "Error number: %d\n", err);
    fprintf(stderr, "Error message: %s\n", Pa_GetErrorText(err));
    return -1;
}
