-- Severities
INSERT INTO severities (id, name, level, description) VALUES
('f47ac10b-58cc-4372-a567-0e02b2c3d479', 'LOW', 1, 'Незначительные проблемы, рекомендуется исправить'),
('7c9e6679-7425-40de-944b-e07fc1f90ae7', 'MEDIUM', 2, 'Проблемы средней опасности, требуют внимания'),
('6ba7b810-9dad-11d1-80b4-00c04fd430c8', 'HIGH', 3, 'Критические уязвимости, необходимо срочно исправить');

-- Issues
INSERT INTO issues (id, name, description, recommendation, severity_id) VALUES
(
    'a1b2c3d4-e5f6-4789-a123-456789abcdef',
    'Debug mode enabled',
    'Приложение работает в debug-режиме, что приводит к чрезмерному логированию и потенциальной утечке чувствительной информации.',
    'Выключите debug-режим в production. Установите уровень логирования в info или выше.',
    'f47ac10b-58cc-4372-a567-0e02b2c3d479'
),
(
    'bf6c7d8e-9f0a-4b1c-8d9e-0a1b2c3d4e5f',
    'Password in plain text',
    'Пароль указан в открытом виде в конфигурации. Это позволяет любому, кто имеет доступ к файлу, прочитать пароль.',
    'Используйте переменные окружения или секретные менеджеры. Никогда не храните пароли в конфигах в открытом виде.',
    '6ba7b810-9dad-11d1-80b4-00c04fd430c8'
),
(
    'c1d2e3f4-5a6b-4c7d-8e9f-0a1b2c3d4e5f',
    'Binding to 0.0.0.0 without restrictions',
    'Приложение слушает на всех интерфейсах (0.0.0.0), что может сделать его доступным извне без ограничений.',
    'Ограничьте привязку к localhost (127.0.0.1) или используйте межсетевой экран. Добавьте аутентификацию.',
    '7c9e6679-7425-40de-944b-e07fc1f90ae7'
),
(
    'd3e4f5a6-7b8c-4d9e-0f1a-2b3c4d5e6f7a',
    'TLS disabled',
    'TLS проверка отключена (например, insecure-skip-verify), что делает соединение уязвимым для MITM-атак.',
    'Включите проверку TLS. Убедитесь, что используются валидные сертификаты.',
    '6ba7b810-9dad-11d1-80b4-00c04fd430c8'
),
(
    'e5f6a7b8-9c0d-4e1f-2a3b-4c5d6e7f8a9b',
    'Weak hash algorithm',
    'Используется устаревший или небезопасный алгоритм хеширования (например, MD5, SHA1).',
    'Замените на современный алгоритм (SHA-256, bcrypt, Argon2).',
    '6ba7b810-9dad-11d1-80b4-00c04fd430c8'
);

-- Rules
INSERT INTO rules (id, name, description, issue_id, expression, enabled) VALUES
(
    'f1a2b3c4-d5e6-47f8-9a0b-1c2d3e4f5a6b',
    'Debug mode check',
    'Проверяет, включён ли debug-режим в конфигурации',
    'a1b2c3d4-e5f6-4789-a123-456789abcdef',
    'config.log.level == "debug" || config.debug == true',
    true
),
(
    'b2c3d4e5-f6a7-48b9-0c1d-2e3f4a5b6c7d',
    'Plain text password check',
    'Проверяет наличие пароля в открытом виде',
    'bf6c7d8e-9f0a-4b1c-8d9e-0a1b2c3d4e5f',
    '(has(config.password) && config.password != "" && !config.password.matches("^\\$\\{.*\\}$") && !config.password.matches("^\\$[A-Z_]+$")) ||
     (has(config.database) && has(config.database.password) && config.database.password != "" && !config.database.password.matches("^\\$\\{.*\\}$") && !config.database.password.matches("^\\$[A-Z_]+$")) ||
     (has(config.api_key) && config.api_key != "" && !config.api_key.matches("^\\$\\{.*\\}$") && !config.api_key.matches("^\\$[A-Z_]+$"))',
    true
),
(
    'c3d4e5f6-a7b8-49c0-1d2e-3f4a5b6c7d8e',
    'Binding to 0.0.0.0 check',
    'Проверяет, слушает ли приложение на 0.0.0.0 без ограничений',
    'c1d2e3f4-5a6b-4c7d-8e9f-0a1b2c3d4e5f',
    'has(config.server) && has(config.server.host) && config.server.host == "0.0.0.0" ||
     has(config.host) && config.host == "0.0.0.0" ||
     has(config.bind) && config.bind == "0.0.0.0"',
    true
),
(
    'd4e5f6a7-b8c9-40d1-2e3f-4a5b6c7d8e9f',
    'TLS disabled check',
    'Проверяет, отключена ли проверка TLS',
    'd3e4f5a6-7b8c-4d9e-0f1a-2b3c4d5e6f7a',
    'has(config.tls) && has(config.tls.insecure_skip_verify) && config.tls.insecure_skip_verify == true ||
     has(config.tls_insecure_skip_verify) && config.tls_insecure_skip_verify == true ||
     has(config.tls) && has(config.tls.disable) && config.tls.disable == true',
    true
),
(
    'e5f6a7b8-c9d0-41e2-3f4a-5b6c7d8e9f0a',
    'Weak hash algorithm check',
    'Проверяет использование слабых алгоритмов хеширования',
    'e5f6a7b8-9c0d-4e1f-2a3b-4c5d6e7f8a9b',
    'config.digest_algorithm in ["MD5", "SHA1", "md5", "sha1"] || (has(config.storage) && has(config.storage.digest_algorithm) && config.storage.digest_algorithm in ["MD5", "SHA1", "md5", "sha1"]) || (has(config.security) && has(config.security.digest_algorithm) && config.security.digest_algorithm in ["MD5", "SHA1", "md5", "sha1"]) || (has(config.crypto) && has(config.crypto.digest_algorithm) && config.crypto.digest_algorithm in ["MD5", "SHA1", "md5", "sha1"])',
    true
);