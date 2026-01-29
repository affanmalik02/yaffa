def test_main_output(capsys):
    from src.main import main
    main()
    captured = capsys.readouterr()
    assert "YAFFA starter" in captured.out
