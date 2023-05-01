import {useEffect, useState} from 'react';
import './App.css';
import {ChooseDirectory, ChooseFile, Convert, GetHomeDir} from "../wailsjs/go/video/FileManager.js"
import {EventsOn, EventsOff} from "../wailsjs/runtime"

const PROGRESS_EVENT = "progress"

function App() : JSX.Element {
    const [filePath, setFilePath] = useState<string>("")
    const [outputDir, setOutputDir] = useState<string>("")
    const [isConverting, setIsConverting] = useState<boolean>(false)
    const [progress, setProgress] = useState<number>(0)

    function init(){
        GetHomeDir().then((res) => {
            setOutputDir(res)
        })
    }

    function handleFileSelect(){
        ChooseFile().then((path) => {
            setFilePath(path)
        })
    }

    function handleSelectOutputDir(){
        ChooseDirectory().then((path) => {
            setOutputDir(path)
        })
    }

    function handleConvert(){
        setIsConverting(true)
        Convert().then((res) => {
            setFilePath(res)
            setIsConverting(false)
        }).catch((err) => {
            setIsConverting(false)
        })
    }

    function registerProgressEvent(){
        EventsOn(PROGRESS_EVENT, (data) => {
            setProgress(data)
            if (data == 100){
                setIsConverting(false)
            }
        })
    }

    useEffect(() => {
        registerProgressEvent()
        init()
        return () => {
            EventsOff(PROGRESS_EVENT)
        }
    }, [])

    return (
        <>
            <div className="main-content">
                <h1>
                    <span>Shd-01</span>
                    <span>Video</span>
                    <span>Converter</span>
                </h1>
                <div className="main-body">
                    <div id="source-container" className="input-selector-container">
                        <label>Origin file</label>
                        <div className="input-group">
                            <input
                                value={filePath}
                                onChange={(e) =>
                                    setFilePath(e.target.value)}
                                id="input-source-path"
                                type="text"/>
                            <button
                                id="btn-source-choose-file"
                                onClick={handleFileSelect}>
                                Open file
                            </button>
                        </div>
                    </div>
                    <div id="output-container" className="input-selector-container">
                        <label>Output directory</label>
                        <div className="input-group">
                            <input
                                value={outputDir}
                                onChange={(e) =>
                                    setOutputDir(e.target.value)}
                                id="input-output-path"
                                type="text"/>
                            <button
                                id="btn-convert-file"
                                disabled={isConverting}
                                onClick={handleSelectOutputDir}>
                                Select Directory
                            </button>
                        </div>
                    </div>
                    <button onClick={handleConvert}  id="btn-start-convert">START</button>
                    <div className="progress-bar" style={{ width: `${progress}%` }} />
                </div>
            </div>
        </>
    )
}

export default App
